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
import type { FixedChargeType } from './FixedChargeType';
import type { InstanceLink } from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import type { WarningType } from './WarningType';
/**
 * Method to update fixed charges of a reservation.
 * @export
 * @interface FixedCharges
 */
export interface FixedCharges {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof FixedCharges
     */
    hotelId?: string;
    /**
     *
     * @type {UniqueIDType}
     * @memberof FixedCharges
     */
    reservationId?: UniqueIDType;
    /**
     * Holds fixed charge detail.
     * @type {Array<FixedChargeType>}
     * @memberof FixedCharges
     */
    fixedCharges?: Array<FixedChargeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof FixedCharges
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FixedCharges
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the FixedCharges interface.
 */
export declare function instanceOfFixedCharges(value: object): boolean;
export declare function FixedChargesFromJSON(json: any): FixedCharges;
export declare function FixedChargesFromJSONTyped(json: any, ignoreDiscriminator: boolean): FixedCharges;
export declare function FixedChargesToJSON(value?: FixedCharges | null): any;
