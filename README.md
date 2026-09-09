# tribepeer (Go)

Go client for TribePeer — e-campus, CBT, AI-authored materials, and per-learner tracking in any Go service.

```bash
go get github.com/devadu001/tribepeer-go
```

## What this unlocks

Same engine as PHP, JavaScript, and Java. You write the handlers. TribePeer runs the classroom.

| Capability | What it does |
|---|---|
| **E-campus / classes** | Faculties, cohorts, join codes, branded portals |
| **Curriculum + e-library** | Modules, lessons, files, scheme of work |
| **Quizzes & CBT** | Generate, sit, auto-grade, results |
| **AI materials** | Lesson notes, assignments, curriculum from a PDF |
| **Ask TribeMate** | Tutor in the product, after hours or between classes |
| **Learner tracking** | Completions, submissions, per-student progress |
| **Cohort chat** | Class threads without a second chat stack |
| **Organisation training** | Staff onboarding on your own site |

The client secret stays on your server.

## Keys

1. [Register](https://www.tribepeer.com/register)
2. [Become a Tribe Owner](https://www.tribepeer.com/tribe-owner/apply)
3. [Issue keys](https://www.tribepeer.com/tribe-owner/credentials)
4. [API guide](https://www.tribepeer.com/institutions/docs)

```env
TP_CLIENT_ID=tp_id_…
TP_CLIENT_SECRET=tp_sec_…
```

```go
import tribepeer "github.com/devadu001/tribepeer-go"

client := tribepeer.New(os.Getenv("TP_CLIENT_ID"), os.Getenv("TP_CLIENT_SECRET"))
classes, err := client.Tribes.List(ctx)

notes, _ := client.AI.Material(ctx, map[string]any{"topic": "Photosynthesis", "level": "SS2"})
quiz, _ := client.AI.Quiz(ctx, map[string]any{"material": notes, "questions": 10})

_, _ = client.Campus.SubmitQuiz(ctx, tribe, material, map[string]any{"answers": answers})
roster, _ := client.Campus.Students(ctx, tribe)
_, _ = client.Campus.Complete(ctx, tribe, material)
```

Go 1.22+. MIT.
