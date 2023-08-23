/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Person's name in an external system.
 * @export
 */
export declare const PersonNameTypeType: {
    readonly Primary: "Primary";
    readonly Alternate: "Alternate";
    readonly Incognito: "Incognito";
    readonly External: "External";
    readonly Phonetic: "Phonetic";
};
export type PersonNameTypeType = typeof PersonNameTypeType[keyof typeof PersonNameTypeType];
export declare function PersonNameTypeTypeFromJSON(json: any): PersonNameTypeType;
export declare function PersonNameTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PersonNameTypeType;
export declare function PersonNameTypeTypeToJSON(value?: PersonNameTypeType | null): any;
