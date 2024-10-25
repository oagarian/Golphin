package response

type Gemini struct {
	Content string `json:"content"`
}

func (g *Gemini) SetContent(content string) error { 
	g.Content = content
    return nil
}

func NewGemini() Gemini {
    return Gemini{}
}
