/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ResProfileTypeType } from './ResProfileTypeType';
import {
    ResProfileTypeTypeFromJSON,
    ResProfileTypeTypeFromJSONTyped,
    ResProfileTypeTypeToJSON,
} from './ResProfileTypeType';

/**
 * Instruction to attach Pre-Configured Routing Instructions to a Reservation.
 * @export
 * @interface ResPreConfiguredRoutingInstrType
 */
export interface ResPreConfiguredRoutingInstrType {
    /**
     * Rate Code with attached Routing Instruction.
     * @type {string}
     * @memberof ResPreConfiguredRoutingInstrType
     */
    ratePlanCode?: string;
    /**
     * 
     * @type {ResProfileTypeType}
     * @memberof ResPreConfiguredRoutingInstrType
     */
    profileType?: ResProfileTypeType;
    /**
     * Promotion Code with attached Complimentary Routing.
     * @type {string}
     * @memberof ResPreConfiguredRoutingInstrType
     */
    promotionCode?: string;
}

/**
 * Check if a given object implements the ResPreConfiguredRoutingInstrType interface.
 */
export function instanceOfResPreConfiguredRoutingInstrType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResPreConfiguredRoutingInstrTypeFromJSON(json: any): ResPreConfiguredRoutingInstrType {
    return ResPreConfiguredRoutingInstrTypeFromJSONTyped(json, false);
}

export function ResPreConfiguredRoutingInstrTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResPreConfiguredRoutingInstrType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'profileType': !exists(json, 'profileType') ? undefined : ResProfileTypeTypeFromJSON(json['profileType']),
        'promotionCode': !exists(json, 'promotionCode') ? undefined : json['promotionCode'],
    };
}

export function ResPreConfiguredRoutingInstrTypeToJSON(value?: ResPreConfiguredRoutingInstrType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ratePlanCode': value.ratePlanCode,
        'profileType': ResProfileTypeTypeToJSON(value.profileType),
        'promotionCode': value.promotionCode,
    };
}

