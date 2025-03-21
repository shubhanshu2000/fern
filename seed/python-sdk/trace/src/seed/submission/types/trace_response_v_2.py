# This file was auto-generated by Fern from our API Definition.

from ...core.pydantic_utilities import UniversalBaseModel
from ...commons.types.debug_key_value_pairs import DebugKeyValuePairs
from ...commons.types.debug_map_value import DebugMapValue
import typing_extensions
from .submission_id import SubmissionId
from ...core.serialization import FieldMetadata
from .traced_file import TracedFile
import typing
from ...commons.types.debug_variable_value import DebugVariableValue
from .expression_location import ExpressionLocation
from .stack_information import StackInformation
from ...core.pydantic_utilities import IS_PYDANTIC_V2
import pydantic


class TraceResponseV2(UniversalBaseModel):
    submission_id: typing_extensions.Annotated[SubmissionId, FieldMetadata(alias="submissionId")]
    line_number: typing_extensions.Annotated[int, FieldMetadata(alias="lineNumber")]
    file: TracedFile
    return_value: typing_extensions.Annotated[
        typing.Optional[DebugVariableValue], FieldMetadata(alias="returnValue")
    ] = None
    expression_location: typing_extensions.Annotated[
        typing.Optional[ExpressionLocation], FieldMetadata(alias="expressionLocation")
    ] = None
    stack: StackInformation
    stdout: typing.Optional[str] = None

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow", frozen=True)  # type: ignore # Pydantic v2
    else:

        class Config:
            frozen = True
            smart_union = True
            extra = pydantic.Extra.allow
