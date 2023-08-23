/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { EffectiveRateType } from './EffectiveRateType';
import type { HotelReservationType } from './HotelReservationType';
import type { InstanceLink } from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PutReinstateReservationRequest
 */
export interface PutReinstateReservationRequest {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof PutReinstateReservationRequest
     */
    hotelId?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof PutReinstateReservationRequest
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     *
     * @type {number}
     * @memberof PutReinstateReservationRequest
     */
    reservationLockHandle?: number;
    /**
     *
     * @type {HotelReservationType}
     * @memberof PutReinstateReservationRequest
     */
    reservation?: HotelReservationType;
    /**
     * Flag that indicates if room inventory check should be skipped when the reservation is being reinstated.
     * @type {boolean}
     * @memberof PutReinstateReservationRequest
     */
    overrideInventory?: boolean;
    /**
     * Flag that indicates if rate code inventory check should be skipped when the reservation is being reinstated.
     * @type {boolean}
     * @memberof PutReinstateReservationRequest
     */
    overrideRates?: boolean;
    /**
     * Flag that indicates if the check on the housekeeping status for out of service should be skipped.
     * @type {boolean}
     * @memberof PutReinstateReservationRequest
     */
    overrideRoomOutOfService?: boolean;
    /**
     * Flag that indicates if the check on room allocation should be skipped.
     * @type {boolean}
     * @memberof PutReinstateReservationRequest
     */
    overrideRoomAllocation?: boolean;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof PutReinstateReservationRequest
     */
    additionalReservationIdList?: Array<UniqueIDType>;
    /**
     * Collection of effective rate amount per guest on specific dates.
     * @type {Array<EffectiveRateType>}
     * @memberof PutReinstateReservationRequest
     */
    effectiveRates?: Array<EffectiveRateType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PutReinstateReservationRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutReinstateReservationRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PutReinstateReservationRequest interface.
 */
export declare function instanceOfPutReinstateReservationRequest(value: object): boolean;
export declare function PutReinstateReservationRequestFromJSON(json: any): PutReinstateReservationRequest;
export declare function PutReinstateReservationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutReinstateReservationRequest;
export declare function PutReinstateReservationRequestToJSON(value?: PutReinstateReservationRequest | null): any;
