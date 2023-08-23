/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Total Pricing Element either Tax or Service Fee
 * @export
 */
export declare const ChargeTypeType: {
    readonly Tax: "Tax";
    readonly ServiceFee: "ServiceFee";
};
export type ChargeTypeType = typeof ChargeTypeType[keyof typeof ChargeTypeType];
export declare function ChargeTypeTypeFromJSON(json: any): ChargeTypeType;
export declare function ChargeTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChargeTypeType;
export declare function ChargeTypeTypeToJSON(value?: ChargeTypeType | null): any;
