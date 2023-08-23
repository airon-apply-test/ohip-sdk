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
 * The value types of parameter, possible types are Boolean, String, Date, and Number.
 * @export
 */
export declare const ChannelParameterValueType: {
    readonly Boolean: "Boolean";
    readonly String: "String";
    readonly Date: "Date";
    readonly Time: "Time";
    readonly Number: "Number";
    readonly SingleSelectLov: "SingleSelectLOV";
    readonly MultiSelectLov: "MultiSelectLOV";
};
export type ChannelParameterValueType = typeof ChannelParameterValueType[keyof typeof ChannelParameterValueType];
export declare function ChannelParameterValueTypeFromJSON(json: any): ChannelParameterValueType;
export declare function ChannelParameterValueTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelParameterValueType;
export declare function ChannelParameterValueTypeToJSON(value?: ChannelParameterValueType | null): any;
