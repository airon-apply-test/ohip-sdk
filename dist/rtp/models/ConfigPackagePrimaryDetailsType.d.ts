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
 * A Config Package Info type.
 * @export
 * @interface ConfigPackagePrimaryDetailsType
 */
export interface ConfigPackagePrimaryDetailsType {
    /**
     * The description of the package.
     * @type {string}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    description?: string;
    /**
     * The short description of the package.
     * @type {string}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    shortDescription?: string;
    /**
     * The Forecast group package belongs to.
     * @type {string}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    forecastGroup?: string;
    /**
     * Arrangement Code.
     * @type {string}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    arrangementCode?: string;
    /**
     * Indicates the begin sell date of the package.
     * @type {Date}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    beginSellDate?: Date;
    /**
     * Indicates the end sell date of the package.
     * @type {Date}
     * @memberof ConfigPackagePrimaryDetailsType
     */
    endSellDate?: Date;
}
/**
 * Check if a given object implements the ConfigPackagePrimaryDetailsType interface.
 */
export declare function instanceOfConfigPackagePrimaryDetailsType(value: object): boolean;
export declare function ConfigPackagePrimaryDetailsTypeFromJSON(json: any): ConfigPackagePrimaryDetailsType;
export declare function ConfigPackagePrimaryDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigPackagePrimaryDetailsType;
export declare function ConfigPackagePrimaryDetailsTypeToJSON(value?: ConfigPackagePrimaryDetailsType | null): any;
