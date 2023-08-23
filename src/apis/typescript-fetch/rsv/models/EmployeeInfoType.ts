/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AddressInfoType } from './AddressInfoType';
import {
    AddressInfoTypeFromJSON,
    AddressInfoTypeFromJSONTyped,
    AddressInfoTypeToJSON,
} from './AddressInfoType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { EmailInfoType } from './EmailInfoType';
import {
    EmailInfoTypeFromJSON,
    EmailInfoTypeFromJSONTyped,
    EmailInfoTypeToJSON,
} from './EmailInfoType';
import type { PersonNameType } from './PersonNameType';
import {
    PersonNameTypeFromJSON,
    PersonNameTypeFromJSONTyped,
    PersonNameTypeToJSON,
} from './PersonNameType';
import type { TelephoneInfoType } from './TelephoneInfoType';
import {
    TelephoneInfoTypeFromJSON,
    TelephoneInfoTypeFromJSONTyped,
    TelephoneInfoTypeToJSON,
} from './TelephoneInfoType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface EmployeeInfoType
 */
export interface EmployeeInfoType {
    /**
     * 
     * @type {PersonNameType}
     * @memberof EmployeeInfoType
     */
    personName?: PersonNameType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof EmployeeInfoType
     */
    profileId?: UniqueIDType;
    /**
     * 
     * @type {AddressInfoType}
     * @memberof EmployeeInfoType
     */
    addressInfo?: AddressInfoType;
    /**
     * 
     * @type {EmailInfoType}
     * @memberof EmployeeInfoType
     */
    emailInfo?: EmailInfoType;
    /**
     * 
     * @type {TelephoneInfoType}
     * @memberof EmployeeInfoType
     */
    phoneInfo?: TelephoneInfoType;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof EmployeeInfoType
     */
    department?: CodeDescriptionType;
    /**
     * Identifies the gender.
     * @type {string}
     * @memberof EmployeeInfoType
     */
    gender?: EmployeeInfoTypeGenderEnum;
    /**
     * Indicates the date of birth as indicated in the document, in ISO 8601 prescribed format.
     * @type {Date}
     * @memberof EmployeeInfoType
     */
    birthDate?: Date;
    /**
     * Indicates the date of birth as masked.
     * @type {string}
     * @memberof EmployeeInfoType
     */
    birthDateMasked?: string;
}


/**
 * @export
 */
export const EmployeeInfoTypeGenderEnum = {
    Male: 'Male',
    Female: 'Female',
    Unknown: 'Unknown'
} as const;
export type EmployeeInfoTypeGenderEnum = typeof EmployeeInfoTypeGenderEnum[keyof typeof EmployeeInfoTypeGenderEnum];


/**
 * Check if a given object implements the EmployeeInfoType interface.
 */
export function instanceOfEmployeeInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmployeeInfoTypeFromJSON(json: any): EmployeeInfoType {
    return EmployeeInfoTypeFromJSONTyped(json, false);
}

export function EmployeeInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmployeeInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'personName': !exists(json, 'personName') ? undefined : PersonNameTypeFromJSON(json['personName']),
        'profileId': !exists(json, 'profileId') ? undefined : UniqueIDTypeFromJSON(json['profileId']),
        'addressInfo': !exists(json, 'addressInfo') ? undefined : AddressInfoTypeFromJSON(json['addressInfo']),
        'emailInfo': !exists(json, 'emailInfo') ? undefined : EmailInfoTypeFromJSON(json['emailInfo']),
        'phoneInfo': !exists(json, 'phoneInfo') ? undefined : TelephoneInfoTypeFromJSON(json['phoneInfo']),
        'department': !exists(json, 'department') ? undefined : CodeDescriptionTypeFromJSON(json['department']),
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'birthDate': !exists(json, 'birthDate') ? undefined : (new Date(json['birthDate'])),
        'birthDateMasked': !exists(json, 'birthDateMasked') ? undefined : json['birthDateMasked'],
    };
}

export function EmployeeInfoTypeToJSON(value?: EmployeeInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'personName': PersonNameTypeToJSON(value.personName),
        'profileId': UniqueIDTypeToJSON(value.profileId),
        'addressInfo': AddressInfoTypeToJSON(value.addressInfo),
        'emailInfo': EmailInfoTypeToJSON(value.emailInfo),
        'phoneInfo': TelephoneInfoTypeToJSON(value.phoneInfo),
        'department': CodeDescriptionTypeToJSON(value.department),
        'gender': value.gender,
        'birthDate': value.birthDate === undefined ? undefined : (value.birthDate.toISOString().substr(0,10)),
        'birthDateMasked': value.birthDateMasked,
    };
}

