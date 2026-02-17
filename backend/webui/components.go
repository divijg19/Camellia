package webui

import (
	"context"
	"fmt"
	"html"
	"io"

	"github.com/a-h/templ"
)

type classItem struct {
	ID         string
	Title      string
	Instructor string
	Location   string
	Price      float64
}

func sampleClasses() []classItem {
	return []classItem{
		{ID: "123", Title: "Beginner Painting", Instructor: "Jane Doe", Location: "Online", Price: 25},
		{ID: "124", Title: "Python 101", Instructor: "Bob Lee", Location: "Hybrid", Price: 30},
	}
}

func homePage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Camellia</title>
  <script src="https://cdn.tailwindcss.com"></script>
  <script src="https://unpkg.com/htmx.org@1.9.12"></script>
</head>
<body class="bg-slate-50 text-slate-900">
  <main class="mx-auto max-w-4xl p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold">Camellia Classes</h1>
      <p class="mt-2 text-slate-600">Templ + Tailwind + HTMX baseline.</p>
    </header>

    <section class="rounded-lg border bg-white p-4 shadow-sm">
      <div class="mb-4 flex flex-col gap-3 sm:flex-row sm:items-center">
        <input id="user-id" name="userId" class="w-full rounded border px-3 py-2" value="guest-user" />
        <button
          class="rounded bg-slate-900 px-4 py-2 font-medium text-white"
          hx-get="/web/classes"
          hx-target="#class-list"
          hx-swap="innerHTML">
          Load classes
        </button>
      </div>

      <div id="class-list" class="space-y-3">
        <p class="text-sm text-slate-500">Click “Load classes” to fetch available classes.</p>
      </div>

      <div id="booking-result" class="mt-4"></div>
    </section>
  </main>
</body>
</html>`)
		return err
	})
}

func classesFragment() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		for _, class := range sampleClasses() {
			_, err := fmt.Fprintf(
				w,
				`<article class="rounded border p-4">
  <h2 class="text-xl font-semibold">%s</h2>
  <p class="text-sm text-slate-600">Instructor: %s · %s</p>
  <p class="mt-1 font-medium">$%.2f</p>
  <form class="mt-3 flex gap-2"
    hx-post="/web/book"
    hx-target="#booking-result"
    hx-swap="innerHTML"
    hx-include="#user-id">
    <input type="hidden" name="classId" value="%s" />
    <button class="rounded bg-emerald-600 px-3 py-1 text-sm font-medium text-white" type="submit">Book</button>
  </form>
</article>`,
				html.EscapeString(class.Title),
				html.EscapeString(class.Instructor),
				html.EscapeString(class.Location),
				class.Price,
				html.EscapeString(class.ID),
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func bookingResult(userID, classID string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := fmt.Fprintf(
			w,
			`<div class="rounded border border-emerald-300 bg-emerald-50 p-3 text-sm text-emerald-900">
Booking confirmed for user <strong>%s</strong> in class <strong>%s</strong>.
</div>`,
			html.EscapeString(userID),
			html.EscapeString(classID),
		)
		return err
	})
}
