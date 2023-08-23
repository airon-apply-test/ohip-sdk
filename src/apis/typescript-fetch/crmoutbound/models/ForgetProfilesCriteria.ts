/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ForgetProfilesCriteriaType } from './ForgetProfilesCriteriaType';
import {
    ForgetProfilesCriteriaTypeFromJSON,
    ForgetProfilesCriteriaTypeFromJSONTyped,
    ForgetProfilesCriteriaTypeToJSON,
} from './ForgetProfilesCriteriaType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';

/**
 * Operation to validate profiles for forgetting.
 * @export
 * @interface ForgetProfilesCriteria
 */
export interface ForgetProfilesCriteria {
    /**
     * 
     * @type {ForgetProfilesCriteriaType}
     * @memberof ForgetProfilesCriteria
     */
    validateForgetProfilesCriteria?: ForgetProfilesCriteriaType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ForgetProfilesCriteria
     */
    links?: Array<InstanceLink>;
}

/**
 * Check if a given object implements the ForgetProfilesCriteria interface.
 */
export function instanceOfForgetProfilesCriteria(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ForgetProfilesCriteriaFromJSON(json: any): ForgetProfilesCriteria {
    return ForgetProfilesCriteriaFromJSONTyped(json, false);
}

export function ForgetProfilesCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): ForgetProfilesCriteria {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'validateForgetProfilesCriteria': !exists(json, 'validateForgetProfilesCriteria') ? undefined : ForgetProfilesCriteriaTypeFromJSON(json['validateForgetProfilesCriteria']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
    };
}

export function ForgetProfilesCriteriaToJSON(value?: ForgetProfilesCriteria | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'validateForgetProfilesCriteria': ForgetProfilesCriteriaTypeToJSON(value.validateForgetProfilesCriteria),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
    };
}

