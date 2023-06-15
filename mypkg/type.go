package mypkg

type Message struct { // 実際のメッセージ情報
	Role    string `json:"role"`    // 送信者情報(user:人間, assistant:ChatGPT, system:ChatGPTの役情報(キャラ設定みたいなもん))
	Content string `json:"content"` // 送信メッセージ
}
type OpenaiRequest struct {
	Model     string    `json:"model"`      // 使用するAPIモデルの設定．参照:https://platform.openai.com/account/rate-limits
	Messages  []Message `json:"messages"`   // これまでの会話を構成するメッセージリスト
	MaxTokens int       `json:"max_tokens"` // 返答メッセージの最大トークン数．MaxTokensを超えるメッセージの場合メッセージがぶつ切りになりChoice.FinishReasonがlengthになって返却される．
}

type Choice struct {
	Index        int     `json:"index"` // 選択肢のインデックス．特に設定しなければ基本0．つまり一つの選択肢しか返さない)
	Messages     Message `json:"message"`
	FinishReason string  `json:"finish_reason"` // "stop"なら完全終了．"length"はMaxTokensの制限による終了
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`     // APIに送信される入力のトークン数
	CompletionTokens int `json:"completion_tokens"` // APIからの応答として受け取る出力トークン数．MaxTokensで上限設定可能
	TotalTokens      int `json:"total_tokens"`      // 合計トークン数: PromptTokens + CompletionTokens
}

type OpenaiResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"` // 返却される情報の選択肢のリスト(複数の場合もあり，その数分トークンも消費する)
	Usages  Usage    `json:"usage"`
}
