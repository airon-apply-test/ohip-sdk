/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Validations type record returned after Validations are done.
 * @export
 * @interface ResGuaranteeType
 */
export interface ResGuaranteeType {
    /**
     * Guarantee Code.
     * @type {string}
     * @memberof ResGuaranteeType
     */
    guaranteeCode?: string;
    /**
     * Guarantee Code.
     * @type {string}
     * @memberof ResGuaranteeType
     */
    shortDescription?: string;
    /**
     * Guarantee Code.
     * @type {boolean}
     * @memberof ResGuaranteeType
     */
    onHold?: boolean;
}

/**
 * Check if a given object implements the ResGuaranteeType interface.
 */
export function instanceOfResGuaranteeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResGuaranteeTypeFromJSON(json: any): ResGuaranteeType {
    return ResGuaranteeTypeFromJSONTyped(json, false);
}

export function ResGuaranteeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResGuaranteeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'guaranteeCode': !exists(json, 'guaranteeCode') ? undefined : json['guaranteeCode'],
        'shortDescription': !exists(json, 'shortDescription') ? undefined : json['shortDescription'],
        'onHold': !exists(json, 'onHold') ? undefined : json['onHold'],
    };
}

export function ResGuaranteeTypeToJSON(value?: ResGuaranteeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'guaranteeCode': value.guaranteeCode,
        'shortDescription': value.shortDescription,
        'onHold': value.onHold,
    };
}

