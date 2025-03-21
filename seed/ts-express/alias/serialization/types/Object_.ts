/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../index";
import * as SeedAlias from "../../api/index";
import * as core from "../../core";

export const Object_: core.serialization.ObjectSchema<serializers.Object_.Raw, SeedAlias.Object_> =
    core.serialization.lazyObject(() => serializers.Type);

export declare namespace Object_ {
    export type Raw = serializers.Type.Raw;
}
