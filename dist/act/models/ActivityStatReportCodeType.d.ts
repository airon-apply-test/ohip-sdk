/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Supported Activity report codes.
 * @export
 */
export declare const ActivityStatReportCodeType: {
    readonly Activities: "Activities";
};
export type ActivityStatReportCodeType = typeof ActivityStatReportCodeType[keyof typeof ActivityStatReportCodeType];
export declare function ActivityStatReportCodeTypeFromJSON(json: any): ActivityStatReportCodeType;
export declare function ActivityStatReportCodeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityStatReportCodeType;
export declare function ActivityStatReportCodeTypeToJSON(value?: ActivityStatReportCodeType | null): any;
