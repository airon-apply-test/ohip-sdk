/**
 * OPERA Cloud Reservation Master Data Management API
 * APIs to cater for Reservation Configuration in OPERA Cloud. In this module you can retrieve, create, modify or delete configuration related to Reservations, Blocks and Leisure Management.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { GuaranteeConfigTypeMassCancellation } from './GuaranteeConfigTypeMassCancellation';
import type { GuaranteeRequirementsType } from './GuaranteeRequirementsType';
import type { TranslationTextType80 } from './TranslationTextType80';
/**
 * Specifies Guarantee Code attributes.
 * @export
 * @interface GuaranteeConfigType
 */
export interface GuaranteeConfigType {
    /**
     * Description of the Guarantee Code.
     * @type {string}
     * @memberof GuaranteeConfigType
     */
    description?: string;
    /**
     *
     * @type {GuaranteeRequirementsType}
     * @memberof GuaranteeConfigType
     */
    requirements?: GuaranteeRequirementsType;
    /**
     *
     * @type {TranslationTextType80}
     * @memberof GuaranteeConfigType
     */
    shortDescription?: TranslationTextType80;
    /**
     * Payment card code like AX,VI etc.
     * @type {Array<string>}
     * @memberof GuaranteeConfigType
     */
    paymentTypes?: Array<string>;
    /**
     * Code assigned to the Guarantee.
     * @type {string}
     * @memberof GuaranteeConfigType
     */
    guaranteeCode?: string;
    /**
     * If true indicates this Guarantee Code is used only to hold the inventory during reservation process.
     * @type {boolean}
     * @memberof GuaranteeConfigType
     */
    onHold?: boolean;
    /**
     * If true indicates inventory will be reserved/deducted when reservation uses this Guarantee Code.
     * @type {boolean}
     * @memberof GuaranteeConfigType
     */
    reserveInventory?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof GuaranteeConfigType
     */
    orderSequence?: number;
    /**
     * Represents the late arrival time.
     * @type {string}
     * @memberof GuaranteeConfigType
     */
    lateArrival?: string;
    /**
     *
     * @type {GuaranteeConfigTypeMassCancellation}
     * @memberof GuaranteeConfigType
     */
    massCancellation?: GuaranteeConfigTypeMassCancellation;
    /**
     * Indicates the reservation type is inactive or not.
     * @type {boolean}
     * @memberof GuaranteeConfigType
     */
    inactive?: boolean;
}
/**
 * Check if a given object implements the GuaranteeConfigType interface.
 */
export declare function instanceOfGuaranteeConfigType(value: object): boolean;
export declare function GuaranteeConfigTypeFromJSON(json: any): GuaranteeConfigType;
export declare function GuaranteeConfigTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuaranteeConfigType;
export declare function GuaranteeConfigTypeToJSON(value?: GuaranteeConfigType | null): any;
