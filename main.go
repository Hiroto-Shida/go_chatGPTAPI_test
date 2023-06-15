package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"golangtest/mypkg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var messages []mypkg.Message

// .env読み込み
func loadApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".envの読み込みが出来ませんでした: %v", err)
	}
	API_KEY := os.Getenv("API_KEY")
	return API_KEY
}

// ChatGPTにリクエスト
func getOpenAIResponse(apiURL, apiKey string) []byte {
	// リクエストボディを作成
	requestData := mypkg.OpenaiRequest{
		Model:     "gpt-3.5-turbo", // モデルの選択
		Messages:  messages,        // メッセージ内容(それまでの対話メッセージ全てを含む)
		MaxTokens: 100,             // 返答メッセージの最大トークン数の制限
	}

	requestBody, _ := json.Marshal(requestData) // structをjsonに変換

	// POSTリクエストを作成
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	// AuthorizationヘッダーにAPIキーを設定
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントを作成し、リクエストを送信
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// レスポンスのボディを読み取る
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body // jsonのまま返却
}

// メイン関数
func main() {
	// ChatGPT APIのエンドポイントURLを指定
	apiURL := "https://api.openai.com/v1/chat/completions"

	// .envのAPIキーを読み込み
	apiKey := loadApiKey()

	// ChatGPTの初期system設定
	// system_content := "発言の語尾には必ずモアをつけてください. モアイに関する質問が来た場合にのみ応答し,それ以外の質問が来た時には「hege」とだけ返答せよ"
	system_content := "Always end your remarks with 'モア'. Respond only to questions about moai, and only 'hoge' to other questions."
	messages = append(messages, mypkg.Message{ // ユーザの質問をmessagesに追加
		Role:    "system",
		Content: system_content,
	})

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 2; i++ { // 質問→応答を2回繰り返す
		fmt.Print("Ask a question: ")
		question, _ := reader.ReadString('\n') // 標準入力の受付(改行までの文字列を読み取る)
		question = strings.TrimSpace(question) // 入力の空白削除

		if question == "exit" { // exitで終了
			break
		}

		messages = append(messages, mypkg.Message{ // ユーザの質問をmessagesに追加
			Role:    "user",
			Content: question,
		})

		// chatGPTに返答を求める
		response := getOpenAIResponse(apiURL, apiKey)
		fmt.Println(string(response)) // 返答内容を表示

		// エラーハンドリングしつつ，返答のjsonをgoのObjectに変換
		var responseObject mypkg.OpenaiResponse
		if err := json.Unmarshal(response, &responseObject); err != nil {
			panic(err)
		}
		messages = append(messages, mypkg.Message{ // chatGPTの返答をmessagesに追加
			Role:    "assistant",
			Content: responseObject.Choices[0].Messages.Content,
		})

		fmt.Println("----------------------")
	}

}
