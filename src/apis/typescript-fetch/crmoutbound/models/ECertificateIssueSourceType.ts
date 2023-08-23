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


/**
 * Indicates that OPERA E-Certificate is issued by external system.
 * @export
 */
export const ECertificateIssueSourceType = {
    Opera: 'Opera',
    Web: 'Web',
    Interface: 'Interface'
} as const;
export type ECertificateIssueSourceType = typeof ECertificateIssueSourceType[keyof typeof ECertificateIssueSourceType];


export function ECertificateIssueSourceTypeFromJSON(json: any): ECertificateIssueSourceType {
    return ECertificateIssueSourceTypeFromJSONTyped(json, false);
}

export function ECertificateIssueSourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateIssueSourceType {
    return json as ECertificateIssueSourceType;
}

export function ECertificateIssueSourceTypeToJSON(value?: ECertificateIssueSourceType | null): any {
    return value as any;
}

