from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AnswerRequest(_message.Message):
    __slots__ = ("userAnswer", "correctAnswer")
    USERANSWER_FIELD_NUMBER: _ClassVar[int]
    CORRECTANSWER_FIELD_NUMBER: _ClassVar[int]
    userAnswer: str
    correctAnswer: str
    def __init__(self, userAnswer: _Optional[str] = ..., correctAnswer: _Optional[str] = ...) -> None: ...

class AnswerResponse(_message.Message):
    __slots__ = ("isCorrect",)
    ISCORRECT_FIELD_NUMBER: _ClassVar[int]
    isCorrect: bool
    def __init__(self, isCorrect: bool = ...) -> None: ...
