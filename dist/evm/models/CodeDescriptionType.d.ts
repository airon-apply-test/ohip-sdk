/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * This contains a generic code and description information.
 * @export
 * @interface CodeDescriptionType
 */
export interface CodeDescriptionType {
    /**
     * Code.
     * @type {string}
     * @memberof CodeDescriptionType
     */
    code?: string;
    /**
     * description.
     * @type {string}
     * @memberof CodeDescriptionType
     */
    description?: string;
}
/**
 * Check if a given object implements the CodeDescriptionType interface.
 */
export declare function instanceOfCodeDescriptionType(value: object): boolean;
export declare function CodeDescriptionTypeFromJSON(json: any): CodeDescriptionType;
export declare function CodeDescriptionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CodeDescriptionType;
export declare function CodeDescriptionTypeToJSON(value?: CodeDescriptionType | null): any;
