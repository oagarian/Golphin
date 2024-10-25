package request 

type gemini struct {
	prompt string `json:"prompt"`
}

func (g *gemini) GetPrompt() gemini {
    return *g
}