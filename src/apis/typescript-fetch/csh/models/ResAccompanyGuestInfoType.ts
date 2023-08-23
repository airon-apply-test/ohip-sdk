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
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface ResAccompanyGuestInfoType
 */
export interface ResAccompanyGuestInfoType {
    /**
     * Given name, first name or names
     * @type {string}
     * @memberof ResAccompanyGuestInfoType
     */
    firstName?: string;
    /**
     * Family name, last name.
     * @type {string}
     * @memberof ResAccompanyGuestInfoType
     */
    lastName?: string;
    /**
     * String representation of the full name
     * @type {string}
     * @memberof ResAccompanyGuestInfoType
     */
    fullName?: string;
    /**
     * Unique identifier of the police registration card number.
     * @type {string}
     * @memberof ResAccompanyGuestInfoType
     */
    registrationCardNo?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ResAccompanyGuestInfoType
     */
    profileIdList?: Array<UniqueIDType>;
}

/**
 * Check if a given object implements the ResAccompanyGuestInfoType interface.
 */
export function instanceOfResAccompanyGuestInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResAccompanyGuestInfoTypeFromJSON(json: any): ResAccompanyGuestInfoType {
    return ResAccompanyGuestInfoTypeFromJSONTyped(json, false);
}

export function ResAccompanyGuestInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResAccompanyGuestInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'firstName': !exists(json, 'firstName') ? undefined : json['firstName'],
        'lastName': !exists(json, 'lastName') ? undefined : json['lastName'],
        'fullName': !exists(json, 'fullName') ? undefined : json['fullName'],
        'registrationCardNo': !exists(json, 'registrationCardNo') ? undefined : json['registrationCardNo'],
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
    };
}

export function ResAccompanyGuestInfoTypeToJSON(value?: ResAccompanyGuestInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'firstName': value.firstName,
        'lastName': value.lastName,
        'fullName': value.fullName,
        'registrationCardNo': value.registrationCardNo,
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
    };
}

