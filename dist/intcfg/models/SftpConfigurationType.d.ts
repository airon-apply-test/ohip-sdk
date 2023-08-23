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
 * Information which uniquely identifies SFTP Configuration
 * @export
 * @interface SftpConfigurationType
 */
export interface SftpConfigurationType {
    /**
     * Unique id associated with this configuration
     * @type {number}
     * @memberof SftpConfigurationType
     */
    configurationId?: number;
    /**
     * SFTP destination
     * @type {string}
     * @memberof SftpConfigurationType
     */
    destination?: string;
    /**
     * Description of the destination, such as Shift Reports.
     * @type {string}
     * @memberof SftpConfigurationType
     */
    description?: string;
    /**
     * Indicates whether the configuration is inactive or not.
     * @type {boolean}
     * @memberof SftpConfigurationType
     */
    inactive?: boolean;
    /**
     * Hotel code
     * @type {string}
     * @memberof SftpConfigurationType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the SftpConfigurationType interface.
 */
export declare function instanceOfSftpConfigurationType(value: object): boolean;
export declare function SftpConfigurationTypeFromJSON(json: any): SftpConfigurationType;
export declare function SftpConfigurationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SftpConfigurationType;
export declare function SftpConfigurationTypeToJSON(value?: SftpConfigurationType | null): any;
