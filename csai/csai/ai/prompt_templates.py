from langchain_core.prompts import ChatPromptTemplate

def get_prompt(type):
    match type:
        case "analyst":
            return get_analyst_prompt()
        case _:
            return get_basic_prompt()


def get_basic_prompt():
    template = """Question: {question}

    Answer: Let's think step by step."""

    prompt = ChatPromptTemplate.from_template(template)

    return prompt

def get_analyst_prompt():
    template = """
    You are a helpful financial analyst with deep business expertise.  You only
    answer questions based on real knowledge and never hallucinate.

    Question: {question}

    Answer: Let's think step by step."""

    prompt = ChatPromptTemplate.from_template(template)

    return prompt