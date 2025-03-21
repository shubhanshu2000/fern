/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as FernOpenapiIr from "../../../../api/index";
import * as core from "../../../../core";
import { HttpMethod } from "./HttpMethod";
import { TagId } from "../../commons/types/TagId";
import { PathParameter } from "./PathParameter";
import { QueryParameter } from "./QueryParameter";
import { Header } from "./Header";
import { EndpointSdkName } from "./EndpointSdkName";
import { Request } from "./Request";
import { Response } from "./Response";
import { StatusCode } from "../../commons/types/StatusCode";
import { HttpError } from "./HttpError";
import { HttpEndpointServer } from "./HttpEndpointServer";
import { EndpointExample } from "./EndpointExample";
import { Pagination } from "./Pagination";
import { WithDescription } from "../../commons/types/WithDescription";
import { WithAvailability } from "../../commons/types/WithAvailability";
import { WithSource } from "../../commons/types/WithSource";
import { WithNamespace } from "../../commons/types/WithNamespace";

export const Endpoint: core.serialization.ObjectSchema<serializers.Endpoint.Raw, FernOpenapiIr.Endpoint> =
    core.serialization
        .objectWithoutOptionalProperties({
            authed: core.serialization.boolean(),
            internal: core.serialization.boolean().optional(),
            idempotent: core.serialization.boolean().optional(),
            method: HttpMethod,
            audiences: core.serialization.list(core.serialization.string()),
            path: core.serialization.string(),
            summary: core.serialization.string().optional(),
            operationId: core.serialization.string().optional(),
            tags: core.serialization.list(TagId),
            pathParameters: core.serialization.list(PathParameter),
            queryParameters: core.serialization.list(QueryParameter),
            headers: core.serialization.list(Header),
            sdkName: EndpointSdkName.optional(),
            generatedRequestName: core.serialization.string(),
            requestNameOverride: core.serialization.string().optional(),
            request: Request.optional(),
            response: Response.optional(),
            errors: core.serialization.record(StatusCode, HttpError),
            servers: core.serialization.list(HttpEndpointServer),
            examples: core.serialization.list(EndpointExample),
            pagination: Pagination.optional(),
        })
        .extend(WithDescription)
        .extend(WithAvailability)
        .extend(WithSource)
        .extend(WithNamespace);

export declare namespace Endpoint {
    export interface Raw extends WithDescription.Raw, WithAvailability.Raw, WithSource.Raw, WithNamespace.Raw {
        authed: boolean;
        internal?: boolean | null;
        idempotent?: boolean | null;
        method: HttpMethod.Raw;
        audiences: string[];
        path: string;
        summary?: string | null;
        operationId?: string | null;
        tags: TagId.Raw[];
        pathParameters: PathParameter.Raw[];
        queryParameters: QueryParameter.Raw[];
        headers: Header.Raw[];
        sdkName?: EndpointSdkName.Raw | null;
        generatedRequestName: string;
        requestNameOverride?: string | null;
        request?: Request.Raw | null;
        response?: Response.Raw | null;
        errors: Record<StatusCode.Raw, HttpError.Raw>;
        servers: HttpEndpointServer.Raw[];
        examples: EndpointExample.Raw[];
        pagination?: Pagination.Raw | null;
    }
}
