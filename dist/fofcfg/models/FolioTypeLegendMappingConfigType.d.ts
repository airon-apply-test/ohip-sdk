/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Folio type legened mapping information.
 * @export
 * @interface FolioTypeLegendMappingConfigType
 */
export interface FolioTypeLegendMappingConfigType {
    /**
     * Property associated with folio type legend mapping
     * @type {string}
     * @memberof FolioTypeLegendMappingConfigType
     */
    hotelId?: string;
    /**
     * Legend code to which folio types attached.
     * @type {string}
     * @memberof FolioTypeLegendMappingConfigType
     */
    legendCode?: string;
    /**
     * Folio type attached to the legend.
     * @type {string}
     * @memberof FolioTypeLegendMappingConfigType
     */
    folioType?: string;
    /**
     * Folio type description attached to the legend.
     * @type {string}
     * @memberof FolioTypeLegendMappingConfigType
     */
    folioTypeDescription?: string;
}
/**
 * Check if a given object implements the FolioTypeLegendMappingConfigType interface.
 */
export declare function instanceOfFolioTypeLegendMappingConfigType(value: object): boolean;
export declare function FolioTypeLegendMappingConfigTypeFromJSON(json: any): FolioTypeLegendMappingConfigType;
export declare function FolioTypeLegendMappingConfigTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FolioTypeLegendMappingConfigType;
export declare function FolioTypeLegendMappingConfigTypeToJSON(value?: FolioTypeLegendMappingConfigType | null): any;
