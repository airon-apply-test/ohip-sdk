/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { RoutingInstructionsToChangeCriteria } from './RoutingInstructionsToChangeCriteria';
import {
    RoutingInstructionsToChangeCriteriaFromJSON,
    RoutingInstructionsToChangeCriteriaFromJSONTyped,
    RoutingInstructionsToChangeCriteriaToJSON,
} from './RoutingInstructionsToChangeCriteria';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface ChangeRoutingInstructionsRequest
 */
export interface ChangeRoutingInstructionsRequest {
    /**
     * 
     * @type {RoutingInstructionsToChangeCriteria}
     * @memberof ChangeRoutingInstructionsRequest
     */
    criteria?: RoutingInstructionsToChangeCriteria;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ChangeRoutingInstructionsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChangeRoutingInstructionsRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ChangeRoutingInstructionsRequest interface.
 */
export function instanceOfChangeRoutingInstructionsRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ChangeRoutingInstructionsRequestFromJSON(json: any): ChangeRoutingInstructionsRequest {
    return ChangeRoutingInstructionsRequestFromJSONTyped(json, false);
}

export function ChangeRoutingInstructionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeRoutingInstructionsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'criteria': !exists(json, 'criteria') ? undefined : RoutingInstructionsToChangeCriteriaFromJSON(json['criteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ChangeRoutingInstructionsRequestToJSON(value?: ChangeRoutingInstructionsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'criteria': RoutingInstructionsToChangeCriteriaToJSON(value.criteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

