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
import type { ReservationId } from './ReservationId';
import type { TraceResolveType } from './TraceResolveType';
import type { TraceTimeInfoType } from './TraceTimeInfoType';
/**
 * Extended Trace object to hold information for a reservation.
 * @export
 * @interface HotelReservationTraceType
 */
export interface HotelReservationTraceType {
    /**
     *
     * @type {TraceTimeInfoType}
     * @memberof HotelReservationTraceType
     */
    timeInfo?: TraceTimeInfoType;
    /**
     *
     * @type {ReservationId}
     * @memberof HotelReservationTraceType
     */
    reservationId?: ReservationId;
    /**
     * Indicates the Department code.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    departmentId?: string;
    /**
     * The information this trace contains.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    traceText?: string;
    /**
     *
     * @type {TraceResolveType}
     * @memberof HotelReservationTraceType
     */
    resolveInfo?: TraceResolveType;
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof HotelReservationTraceType
     */
    idExtension?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof HotelReservationTraceType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof HotelReservationTraceType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof HotelReservationTraceType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof HotelReservationTraceType
     */
    purgeDate?: Date;
}
/**
 * Check if a given object implements the HotelReservationTraceType interface.
 */
export declare function instanceOfHotelReservationTraceType(value: object): boolean;
export declare function HotelReservationTraceTypeFromJSON(json: any): HotelReservationTraceType;
export declare function HotelReservationTraceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelReservationTraceType;
export declare function HotelReservationTraceTypeToJSON(value?: HotelReservationTraceType | null): any;
