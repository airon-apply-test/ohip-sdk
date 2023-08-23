/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { UniqueIDType } from './UniqueIDType';
/**
 * Criteria to validate and change payee tax number.
 * @export
 * @interface ChangePayeeTaxNumberCriteriaType
 */
export interface ChangePayeeTaxNumberCriteriaType {
    /**
     * Property code.
     * @type {string}
     * @memberof ChangePayeeTaxNumberCriteriaType
     */
    hotelId?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ChangePayeeTaxNumberCriteriaType
     */
    payeeId?: UniqueIDType;
    /**
     * Payee Tax Number.
     * @type {string}
     * @memberof ChangePayeeTaxNumberCriteriaType
     */
    taxNumber?: string;
    /**
     * Flag to ignore warnings while updating a payee folio tax number.
     * @type {boolean}
     * @memberof ChangePayeeTaxNumberCriteriaType
     */
    ignoreWarnings?: boolean;
    /**
     * The Cashier ID of the Cashier who is currently processing the transaction(s).
     * @type {number}
     * @memberof ChangePayeeTaxNumberCriteriaType
     */
    cashierId?: number;
}
/**
 * Check if a given object implements the ChangePayeeTaxNumberCriteriaType interface.
 */
export declare function instanceOfChangePayeeTaxNumberCriteriaType(value: object): boolean;
export declare function ChangePayeeTaxNumberCriteriaTypeFromJSON(json: any): ChangePayeeTaxNumberCriteriaType;
export declare function ChangePayeeTaxNumberCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangePayeeTaxNumberCriteriaType;
export declare function ChangePayeeTaxNumberCriteriaTypeToJSON(value?: ChangePayeeTaxNumberCriteriaType | null): any;
