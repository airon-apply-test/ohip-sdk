/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Date of birth field of Membership Type.
 * @export
 */
export const UserDefinedEnrollmentFieldNameType = {
    Communication: 'Communication',
    LastName: 'LastName',
    FirstName: 'FirstName',
    Address: 'Address',
    Fee: 'Fee',
    BirthDate: 'BirthDate'
} as const;
export type UserDefinedEnrollmentFieldNameType = typeof UserDefinedEnrollmentFieldNameType[keyof typeof UserDefinedEnrollmentFieldNameType];


export function UserDefinedEnrollmentFieldNameTypeFromJSON(json: any): UserDefinedEnrollmentFieldNameType {
    return UserDefinedEnrollmentFieldNameTypeFromJSONTyped(json, false);
}

export function UserDefinedEnrollmentFieldNameTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserDefinedEnrollmentFieldNameType {
    return json as UserDefinedEnrollmentFieldNameType;
}

export function UserDefinedEnrollmentFieldNameTypeToJSON(value?: UserDefinedEnrollmentFieldNameType | null): any {
    return value as any;
}

