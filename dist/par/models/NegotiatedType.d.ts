/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { NegotiatedInfoType } from './NegotiatedInfoType';
/**
 * This holds a list of NegotiatedInfoType.
 * @export
 * @interface NegotiatedType
 */
export interface NegotiatedType {
    /**
     *
     * @type {Array<NegotiatedInfoType>}
     * @memberof NegotiatedType
     */
    negotiatedInfoList?: Array<NegotiatedInfoType>;
    /**
     * Hotel code for the negotiated rate.
     * @type {string}
     * @memberof NegotiatedType
     */
    hotelId?: string;
    /**
     * Rate plan code for the negotiated rate.
     * @type {string}
     * @memberof NegotiatedType
     */
    rateCode?: string;
}
/**
 * Check if a given object implements the NegotiatedType interface.
 */
export declare function instanceOfNegotiatedType(value: object): boolean;
export declare function NegotiatedTypeFromJSON(json: any): NegotiatedType;
export declare function NegotiatedTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NegotiatedType;
export declare function NegotiatedTypeToJSON(value?: NegotiatedType | null): any;
