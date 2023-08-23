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


/**
 * 
 * @export
 */
export const ResProfileTypeType = {
    Guest: 'Guest',
    Company: 'Company',
    Group: 'Group',
    TravelAgent: 'TravelAgent',
    Source: 'Source',
    ReservationContact: 'ReservationContact',
    BillingContact: 'BillingContact',
    Addressee: 'Addressee'
} as const;
export type ResProfileTypeType = typeof ResProfileTypeType[keyof typeof ResProfileTypeType];


export function ResProfileTypeTypeFromJSON(json: any): ResProfileTypeType {
    return ResProfileTypeTypeFromJSONTyped(json, false);
}

export function ResProfileTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResProfileTypeType {
    return json as ResProfileTypeType;
}

export function ResProfileTypeTypeToJSON(value?: ResProfileTypeType | null): any {
    return value as any;
}

