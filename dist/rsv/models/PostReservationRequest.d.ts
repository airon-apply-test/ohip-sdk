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
import type { ChannelResvRQInfoType } from './ChannelResvRQInfoType';
import type { HotelReservationsType } from './HotelReservationsType';
import type { InstanceLink } from './InstanceLink';
import type { ReservationInstructionType } from './ReservationInstructionType';
import type { ReservationsInstructionsType } from './ReservationsInstructionsType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostReservationRequest
 */
export interface PostReservationRequest {
    /**
     *
     * @type {HotelReservationsType}
     * @memberof PostReservationRequest
     */
    reservations?: HotelReservationsType;
    /**
     * Instruction on what has to be fetched. Refer to Generic common types document.
     * @type {Array<ReservationInstructionType>}
     * @memberof PostReservationRequest
     */
    fetchInstructions?: Array<ReservationInstructionType>;
    /**
     *
     * @type {ReservationsInstructionsType}
     * @memberof PostReservationRequest
     */
    reservationsInstructionsType?: ReservationsInstructionsType;
    /**
     *
     * @type {ChannelResvRQInfoType}
     * @memberof PostReservationRequest
     */
    channelInformation?: ChannelResvRQInfoType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostReservationRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostReservationRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostReservationRequest interface.
 */
export declare function instanceOfPostReservationRequest(value: object): boolean;
export declare function PostReservationRequestFromJSON(json: any): PostReservationRequest;
export declare function PostReservationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostReservationRequest;
export declare function PostReservationRequestToJSON(value?: PostReservationRequest | null): any;
