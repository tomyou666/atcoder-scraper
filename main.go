package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ProblemData struct {
	Problem     string   `json:"problem"`
	Constraints string   `json:"constraints"`
	Input       string   `json:"input"`
	Images      []string `json:"images,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "使用方法: %s <url> [ディレクトリ名/ファイル名]\n", os.Args[0])
		os.Exit(1)
	}

	problemURL := os.Args[1]

	// 問題データを取得
	problemData, err := fetchProblemData(problemURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	// JSONに変換
	jsonData, err := json.MarshalIndent(problemData, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: JSON変換に失敗しました: %v\n", err)
		os.Exit(1)
	}

	// 出力先を決定
	if len(os.Args) >= 3 {
		outputPath := os.Args[2]

		// ディレクトリかファイルかを判定（拡張子があるかどうか）
		if filepath.Ext(outputPath) == "" {
			// ディレクトリとして扱う
			err := os.MkdirAll(outputPath, 0755)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: ディレクトリの作成に失敗しました: %v\n", err)
				os.Exit(1)
			}

			// JSONファイルを保存
			jsonPath := filepath.Join(outputPath, "problem.json")
			err = writeToFile(jsonPath, string(jsonData))
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: ファイルへの書き込みに失敗しました: %v\n", err)
				os.Exit(1)
			}

			// 画像をダウンロード
			if len(problemData.Images) > 0 {
				err = downloadImages(problemURL, problemData.Images, outputPath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "警告: 画像のダウンロードに失敗しました: %v\n", err)
				}
			}

			fmt.Printf("問題データを %s に保存しました\n", outputPath)
		} else {
			// ファイルとして扱う
			err := writeToFile(outputPath, string(jsonData))
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: ファイルへの書き込みに失敗しました: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("問題データを %s に保存しました\n", outputPath)
		}
	} else {
		// 標準出力
		fmt.Print(string(jsonData))
	}
}

func fetchProblemData(problemURL string) (*ProblemData, error) {
	// HTTPリクエスト
	resp, err := http.Get(problemURL)
	if err != nil {
		return nil, fmt.Errorf("HTTPリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTPステータスコード: %d", resp.StatusCode)
	}

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTMLのパースに失敗しました: %w", err)
	}

	data := &ProblemData{}
	var images []string

	// #task-statement .lang-ja セクションを取得
	doc.Find("#task-statement .lang-ja").Each(func(i int, s *goquery.Selection) {
		// 各セクション（part）を処理
		s.Find(".part").Each(func(j int, part *goquery.Selection) {
			// セクションの見出しを取得
			heading := strings.TrimSpace(part.Find("h3").First().Text())
			// 見出しの後の内容を取得
			content := part.Clone()
			content.Find("h3").First().Remove()
			text := strings.TrimSpace(content.Text())

			// 見出しに応じて分類
			if strings.Contains(heading, "問題") || strings.Contains(heading, "Problem") {
				if data.Problem == "" {
					data.Problem = text
				}
			} else if strings.Contains(heading, "制約") || strings.Contains(heading, "Constraints") {
				if data.Constraints == "" {
					data.Constraints = text
				}
			} else if strings.Contains(heading, "入力") || strings.Contains(heading, "Input") {
				if data.Input == "" {
					data.Input = text
				}
			}

			// 画像を検出
			part.Find("img").Each(func(k int, img *goquery.Selection) {
				src, exists := img.Attr("src")
				if exists && src != "" {
					images = append(images, src)
				}
			})
		})

		// .part がない場合、全体から抽出を試みる
		if data.Problem == "" && data.Constraints == "" && data.Input == "" {
			// 問題文全体を取得
			text := strings.TrimSpace(s.Text())
			if text != "" {
				data.Problem = text
			}
		}

		// 画像を検出（.part の外にもある可能性がある）
		s.Find("img").Each(func(k int, img *goquery.Selection) {
			src, exists := img.Attr("src")
			if exists && src != "" {
				// 重複チェック
				found := false
				for _, existing := range images {
					if existing == src {
						found = true
						break
					}
				}
				if !found {
					images = append(images, src)
				}
			}
		})
	})

	// もし上記で取得できなかった場合、#task-statement 全体から取得
	if data.Problem == "" && data.Constraints == "" && data.Input == "" {
		doc.Find("#task-statement").Each(func(i int, s *goquery.Selection) {
			langJa := s.Find(".lang-ja")
			if langJa.Length() > 0 {
				text := strings.TrimSpace(langJa.Text())
				if text != "" {
					data.Problem = text
				}
			} else {
				text := strings.TrimSpace(s.Text())
				if text != "" {
					data.Problem = text
				}
			}
		})
	}

	if data.Problem == "" && data.Constraints == "" && data.Input == "" {
		return nil, fmt.Errorf("問題文が見つかりませんでした")
	}

	data.Images = images
	return data, nil
}

func downloadImages(baseURL string, imageURLs []string, outputDir string) error {
	base, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("ベースURLのパースに失敗しました: %w", err)
	}

	for i, imgURL := range imageURLs {
		// 相対URLを絶対URLに変換
		parsedURL, err := url.Parse(imgURL)
		if err != nil {
			continue
		}
		absoluteURL := base.ResolveReference(parsedURL).String()

		// 画像をダウンロード
		resp, err := http.Get(absoluteURL)
		if err != nil {
			continue
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			continue
		}

		// ファイル名を決定
		filename := filepath.Base(parsedURL.Path)
		if filename == "" || filename == "." || filename == "/" {
			filename = fmt.Sprintf("image_%d.png", i+1)
		}

		// ファイルに保存
		filePath := filepath.Join(outputDir, filename)
		file, err := os.Create(filePath)
		if err != nil {
			resp.Body.Close()
			continue
		}

		_, err = io.Copy(file, resp.Body)
		file.Close()
		resp.Body.Close()
		if err != nil {
			continue
		}
	}

	return nil
}

func writeToFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}
