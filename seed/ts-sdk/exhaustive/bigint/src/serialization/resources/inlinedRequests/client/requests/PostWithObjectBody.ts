/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../../index";
import * as SeedExhaustive from "../../../../../api/index";
import * as core from "../../../../../core";
import { ObjectWithOptionalField } from "../../../types/resources/object/types/ObjectWithOptionalField";

export const PostWithObjectBody: core.serialization.Schema<
    serializers.PostWithObjectBody.Raw,
    SeedExhaustive.PostWithObjectBody
> = core.serialization.object({
    string: core.serialization.string(),
    integer: core.serialization.number(),
    nestedObject: core.serialization.property("NestedObject", ObjectWithOptionalField),
});

export declare namespace PostWithObjectBody {
    export interface Raw {
        string: string;
        integer: number;
        NestedObject: ObjectWithOptionalField.Raw;
    }
}
