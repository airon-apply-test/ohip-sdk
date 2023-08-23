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
import type { ErrorType } from './ErrorType';
import type { ReservationInfoType } from './ReservationInfoType';
/**
 * Information about the block reservation being changed.
 * @export
 * @interface ChangeBlockReservationType
 */
export interface ChangeBlockReservationType {
    /**
     *
     * @type {ReservationInfoType}
     * @memberof ChangeBlockReservationType
     */
    reservationInfo?: ReservationInfoType;
    /**
     * Returning an empty element of this type indicates the successful processing of an message. This is used in conjunction with the Warning Type to report any warnings or business errors.
     * @type {object}
     * @memberof ChangeBlockReservationType
     */
    success?: object;
    /**
     * An error that occurred during the processing of a message.
     * @type {Array<ErrorType>}
     * @memberof ChangeBlockReservationType
     */
    errors?: Array<ErrorType>;
}
/**
 * Check if a given object implements the ChangeBlockReservationType interface.
 */
export declare function instanceOfChangeBlockReservationType(value: object): boolean;
export declare function ChangeBlockReservationTypeFromJSON(json: any): ChangeBlockReservationType;
export declare function ChangeBlockReservationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeBlockReservationType;
export declare function ChangeBlockReservationTypeToJSON(value?: ChangeBlockReservationType | null): any;
