from langchain_core.prompts import ChatPromptTemplate
from langchain_ollama import OllamaLLM
from csai.ai.prompt_templates import get_prompt

from pydantic import BaseModel


# https://api.python.langchain.com/en/latest/llms/langchain_community.llms.ollama.Ollama.html
llm = OllamaLLM(
    #base_url="http://ollama.ollama.svc:11434",
    model="mistral",
    timeout=300,
)

class Chat(BaseModel):
    userAsk: str
    response: str
    context: str

def ask(request:Chat) -> Chat:
    prompt = get_prompt(request.context)
    chain = prompt | llm

    print(prompt)
    request.response = chain.invoke({"question": request.userAsk, "context": request.context})
    return request

if __name__ == '__main__':
    userRequest = Chat(userAsk="who are the main characters in final fantasy 7", response="")
    response = ask(userRequest)

    print(response.response)