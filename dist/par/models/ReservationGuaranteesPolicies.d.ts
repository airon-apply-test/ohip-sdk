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
import type { GuaranteePolicyType } from './GuaranteePolicyType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object to fetch reservation guarantees.
 * @export
 * @interface ReservationGuaranteesPolicies
 */
export interface ReservationGuaranteesPolicies {
    /**
     * Guarantee Code information with cancellation penalty and deposit policy information.
     * @type {Array<GuaranteePolicyType>}
     * @memberof ReservationGuaranteesPolicies
     */
    resGuarantees?: Array<GuaranteePolicyType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ReservationGuaranteesPolicies
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReservationGuaranteesPolicies
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ReservationGuaranteesPolicies interface.
 */
export declare function instanceOfReservationGuaranteesPolicies(value: object): boolean;
export declare function ReservationGuaranteesPoliciesFromJSON(json: any): ReservationGuaranteesPolicies;
export declare function ReservationGuaranteesPoliciesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationGuaranteesPolicies;
export declare function ReservationGuaranteesPoliciesToJSON(value?: ReservationGuaranteesPolicies | null): any;
