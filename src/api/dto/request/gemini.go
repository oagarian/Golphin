package request 

type Gemini struct {
	Prompt string `json:"prompt"`
}

func (g *Gemini) GetPrompt() string {
    return g.Prompt
}