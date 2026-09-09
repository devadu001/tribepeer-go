package tribepeer

import (
	"context"
	"net/http"
)

type Users struct{ c *Client }

func (u *Users) Register(ctx context.Context, body map[string]any) (map[string]any, error) {
	return u.c.request(ctx, http.MethodPost, u.c.ProductBase+"/auth/register", body, "")
}

func (u *Users) Login(ctx context.Context, body map[string]any) (map[string]any, error) {
	data, err := u.c.request(ctx, http.MethodPost, u.c.ProductBase+"/auth/login", body, "")
	if err != nil {
		return nil, err
	}
	if tok, ok := data["access_token"].(string); ok {
		u.c.userToken = tok
	}
	return data, nil
}

func (u *Users) Verify(ctx context.Context, body map[string]any) (map[string]any, error) {
	data, err := u.c.request(ctx, http.MethodPost, u.c.ProductBase+"/auth/verify", body, "")
	if err != nil {
		return nil, err
	}
	if tok, ok := data["access_token"].(string); ok {
		u.c.userToken = tok
	}
	return data, nil
}

func (u *Users) ResendOTP(ctx context.Context, email string) (map[string]any, error) {
	return u.c.request(ctx, http.MethodPost, u.c.ProductBase+"/auth/resend-otp", map[string]any{"email": email}, "")
}

func (u *Users) Refresh(ctx context.Context) (map[string]any, error) {
	data, err := u.c.request(ctx, http.MethodPost, u.c.ProductBase+"/auth/refresh", nil, u.c.userToken)
	if err != nil {
		return nil, err
	}
	if tok, ok := data["access_token"].(string); ok {
		u.c.userToken = tok
	}
	return data, nil
}

func (u *Users) Me(ctx context.Context) (map[string]any, error) {
	return u.c.request(ctx, http.MethodGet, u.c.ProductBase+"/me", nil, u.c.userToken)
}

type Tribes struct{ c *Client }

func (t *Tribes) List(ctx context.Context) (map[string]any, error) {
	return t.c.partner(ctx, http.MethodGet, "/tribes", nil)
}

func (t *Tribes) Create(ctx context.Context, body map[string]any) (map[string]any, error) {
	return t.c.partner(ctx, http.MethodPost, "/tribes", body)
}

func (t *Tribes) Get(ctx context.Context, uuid string) (map[string]any, error) {
	return t.c.partner(ctx, http.MethodGet, "/tribes/"+uuid, nil)
}

func (t *Tribes) Update(ctx context.Context, uuid string, body map[string]any) (map[string]any, error) {
	return t.c.partner(ctx, http.MethodPatch, "/tribes/"+uuid, body)
}

type Materials struct{ c *Client }

func (m *Materials) List(ctx context.Context, tribeUUID string) (map[string]any, error) {
	return m.c.partner(ctx, http.MethodGet, "/tribes/"+tribeUUID+"/materials", nil)
}

type Submissions struct{ c *Client }

func (s *Submissions) Grade(ctx context.Context, id string, body map[string]any) (map[string]any, error) {
	return s.c.partner(ctx, http.MethodPost, "/submissions/"+id+"/grade", body)
}

type Communities struct{ c *Client }

func (c *Communities) Threads(ctx context.Context, communityUUID string) (map[string]any, error) {
	return c.c.partner(ctx, http.MethodGet, "/communities/"+communityUUID+"/threads", nil)
}

type AI struct{ c *Client }

func (a *AI) Chat(ctx context.Context, body map[string]any) (map[string]any, error) {
	return a.c.partner(ctx, http.MethodPost, "/ai/chat", body)
}

func (a *AI) Usage(ctx context.Context) (map[string]any, error) {
	return a.c.partner(ctx, http.MethodGet, "/ai/usage", nil)
}

type Campus struct{ c *Client }

func (k *Campus) Branding(ctx context.Context, key, institution string) (map[string]any, error) {
	if institution == "" {
		institution = k.c.InstitutionUUID
	}
	href := withQuery(k.c.ProductBase+"/branding", map[string]string{"key": key, "institution": institution})
	return k.c.request(ctx, http.MethodGet, href, nil, "")
}

func (k *Campus) Join(ctx context.Context, joinCode string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/join", map[string]any{"join_code": joinCode})
}

func (k *Campus) TribesMine(ctx context.Context) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/tribes/mine", nil)
}

func (k *Campus) Tribe(ctx context.Context, uuid string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/tribes/"+uuid, nil)
}

func (k *Campus) Complete(ctx context.Context, tribeUUID, materialUUID string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/tribes/"+tribeUUID+"/materials/"+materialUUID+"/complete", nil)
}

func (k *Campus) SubmitQuiz(ctx context.Context, tribeUUID, materialUUID string, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/tribes/"+tribeUUID+"/materials/"+materialUUID+"/quiz", body)
}

func (k *Campus) ChatThreads(ctx context.Context) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/chat/threads", nil)
}

func (k *Campus) ChatMessages(ctx context.Context, thread string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/chat/threads/"+thread+"/messages", nil)
}

func (k *Campus) ChatSend(ctx context.Context, thread, body string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/chat/threads/"+thread+"/messages", map[string]any{"body": body})
}

func (k *Campus) Ask(ctx context.Context, message string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/ai/ask", map[string]any{"message": message})
}

func (k *Campus) ManageTribes(ctx context.Context) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/manage/tribes", nil)
}

func (k *Campus) CreateTribe(ctx context.Context, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/manage/tribes", body)
}

func (k *Campus) UpdateTribe(ctx context.Context, tribeUUID string, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPatch, "/manage/tribes/"+tribeUUID, body)
}

func (k *Campus) Students(ctx context.Context, tribeUUID string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/manage/tribes/"+tribeUUID+"/students", nil)
}

func (k *Campus) JoinCodes(ctx context.Context) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/manage/join-codes", nil)
}

func (k *Campus) CreateJoinCode(ctx context.Context, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/manage/join-codes", body)
}

func (k *Campus) Curriculum(ctx context.Context, tribeUUID string) (map[string]any, error) {
	return k.c.product(ctx, http.MethodGet, "/manage/tribes/"+tribeUUID+"/curriculum", nil)
}

func (k *Campus) CreateModule(ctx context.Context, tribeUUID string, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/manage/tribes/"+tribeUUID+"/modules", body)
}

func (k *Campus) CreateMaterial(ctx context.Context, tribeUUID string, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPost, "/manage/tribes/"+tribeUUID+"/materials", body)
}

func (k *Campus) UpdateMaterial(ctx context.Context, tribeUUID, materialUUID string, body map[string]any) (map[string]any, error) {
	return k.c.product(ctx, http.MethodPatch, "/manage/tribes/"+tribeUUID+"/materials/"+materialUUID, body)
}
