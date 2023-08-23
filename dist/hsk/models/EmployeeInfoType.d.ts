/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AddressInfoType } from './AddressInfoType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { EmailInfoType } from './EmailInfoType';
import type { PersonNameType } from './PersonNameType';
import type { TelephoneInfoType } from './TelephoneInfoType';
import type { UniqueIDType } from './UniqueIDType';
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
export declare const EmployeeInfoTypeGenderEnum: {
    readonly Male: "Male";
    readonly Female: "Female";
    readonly Unknown: "Unknown";
};
export type EmployeeInfoTypeGenderEnum = typeof EmployeeInfoTypeGenderEnum[keyof typeof EmployeeInfoTypeGenderEnum];
/**
 * Check if a given object implements the EmployeeInfoType interface.
 */
export declare function instanceOfEmployeeInfoType(value: object): boolean;
export declare function EmployeeInfoTypeFromJSON(json: any): EmployeeInfoType;
export declare function EmployeeInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmployeeInfoType;
export declare function EmployeeInfoTypeToJSON(value?: EmployeeInfoType | null): any;
