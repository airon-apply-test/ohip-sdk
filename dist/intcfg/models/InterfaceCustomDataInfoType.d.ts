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
import type { InterfaceCustomDataDetailType } from './InterfaceCustomDataDetailType';
/**
 * This type holds the custom data of a hotel interface.
 * @export
 * @interface InterfaceCustomDataInfoType
 */
export interface InterfaceCustomDataInfoType {
    /**
     * Specifies whether the details have to be exported as XML. If true, then the details are exported as XML else details are exported as text.
     * @type {boolean}
     * @memberof InterfaceCustomDataInfoType
     */
    exportAsXml?: boolean;
    /**
     * Specifies whether the Doorcard field details have to be included in the XML Export. If true, Doorcard field details will be included in the XML Export else details will not be included. This field is available only when the Export as XML field is selected.
     * @type {boolean}
     * @memberof InterfaceCustomDataInfoType
     */
    includeDoorcardField?: boolean;
    /**
     * Collection of custom data details of a hotel interface.
     * @type {Array<InterfaceCustomDataDetailType>}
     * @memberof InterfaceCustomDataInfoType
     */
    interfaceCustomDataDetails?: Array<InterfaceCustomDataDetailType>;
}
/**
 * Check if a given object implements the InterfaceCustomDataInfoType interface.
 */
export declare function instanceOfInterfaceCustomDataInfoType(value: object): boolean;
export declare function InterfaceCustomDataInfoTypeFromJSON(json: any): InterfaceCustomDataInfoType;
export declare function InterfaceCustomDataInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InterfaceCustomDataInfoType;
export declare function InterfaceCustomDataInfoTypeToJSON(value?: InterfaceCustomDataInfoType | null): any;
