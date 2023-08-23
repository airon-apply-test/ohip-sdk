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
 * XML Posting Interface
 * @export
 */
export declare const HotelInterfaceTypeType: {
    readonly Bms: "Bms";
    readonly Cas: "Cas";
    readonly Ccw: "Ccw";
    readonly Dls: "Dls";
    readonly Eft: "Eft";
    readonly Exp: "Exp";
    readonly Mak: "Mak";
    readonly Mbs: "Mbs";
    readonly Msc: "Msc";
    readonly Pbx: "Pbx";
    readonly Pos: "Pos";
    readonly Svs: "Svs";
    readonly Tik: "Tik";
    readonly Vid: "Vid";
    readonly Vms: "Vms";
    readonly Www: "Www";
    readonly Xml: "Xml";
};
export type HotelInterfaceTypeType = typeof HotelInterfaceTypeType[keyof typeof HotelInterfaceTypeType];
export declare function HotelInterfaceTypeTypeFromJSON(json: any): HotelInterfaceTypeType;
export declare function HotelInterfaceTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInterfaceTypeType;
export declare function HotelInterfaceTypeTypeToJSON(value?: HotelInterfaceTypeType | null): any;
