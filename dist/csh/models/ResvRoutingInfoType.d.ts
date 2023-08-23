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
import type { ReservationInfoType } from './ReservationInfoType';
import type { ResvRoutingInfoTypeComp } from './ResvRoutingInfoTypeComp';
import type { ResvRoutingInfoTypeFolio } from './ResvRoutingInfoTypeFolio';
import type { ResvRoutingInfoTypeRequest } from './ResvRoutingInfoTypeRequest';
import type { ResvRoutingInfoTypeRoom } from './ResvRoutingInfoTypeRoom';
/**
 * A routing info object can either be of type Folio OR of type Room with its corresponding object.
 * @export
 * @interface ResvRoutingInfoType
 */
export interface ResvRoutingInfoType {
    /**
     *
     * @type {ResvRoutingInfoTypeFolio}
     * @memberof ResvRoutingInfoType
     */
    folio?: ResvRoutingInfoTypeFolio;
    /**
     *
     * @type {ResvRoutingInfoTypeRoom}
     * @memberof ResvRoutingInfoType
     */
    room?: ResvRoutingInfoTypeRoom;
    /**
     *
     * @type {ResvRoutingInfoTypeComp}
     * @memberof ResvRoutingInfoType
     */
    comp?: ResvRoutingInfoTypeComp;
    /**
     *
     * @type {ResvRoutingInfoTypeRequest}
     * @memberof ResvRoutingInfoType
     */
    request?: ResvRoutingInfoTypeRequest;
    /**
     * On a successful update, the transactions that are already posted in the guest's folio will be re-organized based on the configured instructions.
     * @type {boolean}
     * @memberof ResvRoutingInfoType
     */
    refreshFolio?: boolean;
    /**
     *
     * @type {ReservationInfoType}
     * @memberof ResvRoutingInfoType
     */
    reservationInfo?: ReservationInfoType;
}
/**
 * Check if a given object implements the ResvRoutingInfoType interface.
 */
export declare function instanceOfResvRoutingInfoType(value: object): boolean;
export declare function ResvRoutingInfoTypeFromJSON(json: any): ResvRoutingInfoType;
export declare function ResvRoutingInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResvRoutingInfoType;
export declare function ResvRoutingInfoTypeToJSON(value?: ResvRoutingInfoType | null): any;
