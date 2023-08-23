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
import type { BlockRateOverrideReasonType } from './BlockRateOverrideReasonType';
import {
    BlockRateOverrideReasonTypeFromJSON,
    BlockRateOverrideReasonTypeFromJSONTyped,
    BlockRateOverrideReasonTypeToJSON,
} from './BlockRateOverrideReasonType';
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
 * Request object for creating Block Rate Override Reasons.
 * @export
 * @interface BlockRateOverrideReasonsCriteria
 */
export interface BlockRateOverrideReasonsCriteria {
    /**
     * List of Block Rate Override Reasons.
     * @type {Array<BlockRateOverrideReasonType>}
     * @memberof BlockRateOverrideReasonsCriteria
     */
    blockRateOverrideReasons?: Array<BlockRateOverrideReasonType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof BlockRateOverrideReasonsCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BlockRateOverrideReasonsCriteria
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the BlockRateOverrideReasonsCriteria interface.
 */
export function instanceOfBlockRateOverrideReasonsCriteria(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockRateOverrideReasonsCriteriaFromJSON(json: any): BlockRateOverrideReasonsCriteria {
    return BlockRateOverrideReasonsCriteriaFromJSONTyped(json, false);
}

export function BlockRateOverrideReasonsCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockRateOverrideReasonsCriteria {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockRateOverrideReasons': !exists(json, 'blockRateOverrideReasons') ? undefined : ((json['blockRateOverrideReasons'] as Array<any>).map(BlockRateOverrideReasonTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function BlockRateOverrideReasonsCriteriaToJSON(value?: BlockRateOverrideReasonsCriteria | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockRateOverrideReasons': value.blockRateOverrideReasons === undefined ? undefined : ((value.blockRateOverrideReasons as Array<any>).map(BlockRateOverrideReasonTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

