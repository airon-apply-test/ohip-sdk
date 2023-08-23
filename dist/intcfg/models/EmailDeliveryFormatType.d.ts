/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Email delivery format
 * @export
 */
export declare const EmailDeliveryFormatType: {
    readonly Htmlcss: "Htmlcss";
    readonly Pdf: "Pdf";
    readonly Rtf: "Rtf";
    readonly Tif: "Tif";
};
export type EmailDeliveryFormatType = typeof EmailDeliveryFormatType[keyof typeof EmailDeliveryFormatType];
export declare function EmailDeliveryFormatTypeFromJSON(json: any): EmailDeliveryFormatType;
export declare function EmailDeliveryFormatTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailDeliveryFormatType;
export declare function EmailDeliveryFormatTypeToJSON(value?: EmailDeliveryFormatType | null): any;
