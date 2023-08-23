/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BlockRefusedReasonType } from './BlockRefusedReasonType';
import {
    BlockRefusedReasonTypeFromJSON,
    BlockRefusedReasonTypeFromJSONTyped,
    BlockRefusedReasonTypeToJSON,
} from './BlockRefusedReasonType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PostBlockRefusedReasonspostBloRequest
 */
export interface PostBlockRefusedReasonspostBloRequest {
    /**
     * List of Block Refused Reasons.
     * @type {Array<BlockRefusedReasonType>}
     * @memberof PostBlockRefusedReasonspostBloRequest
     */
    blockRefusedReasons?: Array<BlockRefusedReasonType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostBlockRefusedReasonspostBloRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostBlockRefusedReasonspostBloRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostBlockRefusedReasonspostBloRequest interface.
 */
export function instanceOfPostBlockRefusedReasonspostBloRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostBlockRefusedReasonspostBloRequestFromJSON(json: any): PostBlockRefusedReasonspostBloRequest {
    return PostBlockRefusedReasonspostBloRequestFromJSONTyped(json, false);
}

export function PostBlockRefusedReasonspostBloRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostBlockRefusedReasonspostBloRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockRefusedReasons': !exists(json, 'blockRefusedReasons') ? undefined : ((json['blockRefusedReasons'] as Array<any>).map(BlockRefusedReasonTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostBlockRefusedReasonspostBloRequestToJSON(value?: PostBlockRefusedReasonspostBloRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockRefusedReasons': value.blockRefusedReasons === undefined ? undefined : ((value.blockRefusedReasons as Array<any>).map(BlockRefusedReasonTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

