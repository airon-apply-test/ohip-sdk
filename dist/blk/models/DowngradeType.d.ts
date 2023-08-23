/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Do not downgrade membership when the next downgrade process runs. When the downgrade process runs, the membership will be automatically set to Grace.
 * @export
 */
export declare const DowngradeType: {
    readonly Grace: "Grace";
    readonly Never: "Never";
    readonly Period: "Period";
};
export type DowngradeType = typeof DowngradeType[keyof typeof DowngradeType];
export declare function DowngradeTypeFromJSON(json: any): DowngradeType;
export declare function DowngradeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DowngradeType;
export declare function DowngradeTypeToJSON(value?: DowngradeType | null): any;
