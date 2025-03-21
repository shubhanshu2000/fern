/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernDefinition from "../../../../api/index";
import * as core from "../../../../core";
import { ApiConfigurationSchemaInternal } from "./ApiConfigurationSchemaInternal";

export const NamespacedApiConfigurationSchema: core.serialization.ObjectSchema<
    serializers.NamespacedApiConfigurationSchema.Raw,
    FernDefinition.NamespacedApiConfigurationSchema
> = core.serialization.object({
    namespaces: core.serialization.record(core.serialization.string(), ApiConfigurationSchemaInternal),
});

export declare namespace NamespacedApiConfigurationSchema {
    export interface Raw {
        namespaces: Record<string, ApiConfigurationSchemaInternal.Raw>;
    }
}
