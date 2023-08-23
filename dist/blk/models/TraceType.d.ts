/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
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
 *
 * @export
 * @interface TraceType
 */
export interface TraceType {
    /**
     *
     * @type {TraceTimeInfoType}
     * @memberof TraceType
     */
    timeInfo?: TraceTimeInfoType;
    /**
     *
     * @type {ReservationId}
     * @memberof TraceType
     */
    reservationId?: ReservationId;
    /**
     * Indicates the Department code.
     * @type {string}
     * @memberof TraceType
     */
    departmentId?: string;
    /**
     * The information this trace contains.
     * @type {string}
     * @memberof TraceType
     */
    traceText?: string;
    /**
     *
     * @type {TraceResolveType}
     * @memberof TraceType
     */
    resolveInfo?: TraceResolveType;
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof TraceType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof TraceType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof TraceType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof TraceType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof TraceType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof TraceType
     */
    idExtension?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof TraceType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof TraceType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof TraceType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof TraceType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof TraceType
     */
    purgeDate?: Date;
}
/**
 * Check if a given object implements the TraceType interface.
 */
export declare function instanceOfTraceType(value: object): boolean;
export declare function TraceTypeFromJSON(json: any): TraceType;
export declare function TraceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TraceType;
export declare function TraceTypeToJSON(value?: TraceType | null): any;
