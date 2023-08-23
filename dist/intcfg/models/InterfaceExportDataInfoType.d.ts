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
import type { InterfaceExportDataDetailType } from './InterfaceExportDataDetailType';
/**
 * This type holds the export data of a hotel interface.
 * @export
 * @interface InterfaceExportDataInfoType
 */
export interface InterfaceExportDataInfoType {
    /**
     * Collection of export data details of a hotel interface.
     * @type {Array<InterfaceExportDataDetailType>}
     * @memberof InterfaceExportDataInfoType
     */
    interfaceExportDataDetails?: Array<InterfaceExportDataDetailType>;
    /**
     * Collection of export data details of a hotel interface.
     * @type {Array<InterfaceExportDataDetailType>}
     * @memberof InterfaceExportDataInfoType
     */
    interfaceExportDataMandatoryDetails?: Array<InterfaceExportDataDetailType>;
}
/**
 * Check if a given object implements the InterfaceExportDataInfoType interface.
 */
export declare function instanceOfInterfaceExportDataInfoType(value: object): boolean;
export declare function InterfaceExportDataInfoTypeFromJSON(json: any): InterfaceExportDataInfoType;
export declare function InterfaceExportDataInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceExportDataInfoType;
export declare function InterfaceExportDataInfoTypeToJSON(value?: InterfaceExportDataInfoType | null): any;
