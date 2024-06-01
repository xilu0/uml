from openai import OpenAI

client = OpenAI(
  base_url = "http://localhost:11434/v1/",
  api_key = "$API_KEY_REQUIRED_IF_EXECUTING_OUTSIDE_NGC"
)
msg = "Hello, waht is your name?"
completion = client.chat.completions.create(
  model="llama3:8b",
  messages=[
    {
                "role": "system",
                "content": "You are a helpful assistant."
    },
    {"role":"assistant","content":"my name is Dave"},
    {"role":"user","content":msg}
    ],
  temperature=0.5,
  top_p=1,
  max_tokens=1024,
  stream=True
)

for chunk in completion:
  if chunk.choices[0].delta.content is not None:
    print(chunk.choices[0].delta.content, end="")

