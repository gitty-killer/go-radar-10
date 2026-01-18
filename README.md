# go-radar-10

A small Go tool that computes text statistics for radar.

## Objective
- Provide quick text metrics for radar documents.
- Report top word frequencies for fast inspection.

## Use cases
- Validate radar drafts for repeated terms before review.
- Summarize radar notes when preparing reports.

## Usage
go run main.go data/sample.txt --top 5

## Output
- lines: total line count
- words: total word count
- chars: total character count
- top words: most frequent tokens (case-insensitive)

## Testing
- run `bash scripts/verify.sh`

## Notes
- Only ASCII letters and digits are treated as word characters.
