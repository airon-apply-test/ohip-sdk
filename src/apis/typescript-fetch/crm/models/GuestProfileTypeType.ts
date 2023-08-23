/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * The types of Profile handled by the web service.
 * @export
 */
export const GuestProfileTypeType = {
    Guest: 'Guest',
    Employee: 'Employee',
    Contact: 'Contact'
} as const;
export type GuestProfileTypeType = typeof GuestProfileTypeType[keyof typeof GuestProfileTypeType];


export function GuestProfileTypeTypeFromJSON(json: any): GuestProfileTypeType {
    return GuestProfileTypeTypeFromJSONTyped(json, false);
}

export function GuestProfileTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuestProfileTypeType {
    return json as GuestProfileTypeType;
}

export function GuestProfileTypeTypeToJSON(value?: GuestProfileTypeType | null): any {
    return value as any;
}

