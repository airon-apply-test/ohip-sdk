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
 * This provides information for a profile negotiated rate.
 * @export
 * @interface NegotiatedInfoType
 */
export interface NegotiatedInfoType {
    /**
     * The master identifier for multiple offices/locations under the same company profile. This is optional
     * @type {string}
     * @memberof NegotiatedInfoType
     */
    corporateAgreementId?: string;
    /**
     * Informational purposes only in numeric format.
     * @type {string}
     * @memberof NegotiatedInfoType
     */
    comissionCode?: string;
    /**
     * The sell order.
     * @type {number}
     * @memberof NegotiatedInfoType
     */
    order?: number;
    /**
     * Negotiated Rate is inactive or not
     * @type {boolean}
     * @memberof NegotiatedInfoType
     */
    inactive?: boolean;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof NegotiatedInfoType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof NegotiatedInfoType
     */
    end?: Date;
}
/**
 * Check if a given object implements the NegotiatedInfoType interface.
 */
export declare function instanceOfNegotiatedInfoType(value: object): boolean;
export declare function NegotiatedInfoTypeFromJSON(json: any): NegotiatedInfoType;
export declare function NegotiatedInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NegotiatedInfoType;
export declare function NegotiatedInfoTypeToJSON(value?: NegotiatedInfoType | null): any;
