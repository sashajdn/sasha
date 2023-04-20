from typing import List, Union

from langchain.llms import OpenAI
from langchain.schema import SystemMessage, AIMessage, HumanMessage

LLMMessage = Union[SystemMessage, AIMessage, HumanMessage]

class Config:
    openai_api_key: str
    temperature: float

    def __init__(self, temperature: float = 0.5, openai_api_key: str = ""):
        self.temperature = temperature
        self.openai_api_key = openai_api_key


class LangChainClient:
    _open_api_key: str
    _llm: OpenAI

    def __init__(self, config: Config):
        self._open_api_key: str = config.openai_api_key
        self._llm: OpenAI = OpenAI(client={}, temperature=config.temperature)

    def __call__(self, prompt: str):
        return self._llm(prompt)
