/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { ProfileRoutingInstructionsType } from './ProfileRoutingInstructionsType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * The type contains routing instructions for the profile.
 * @export
 * @interface ProfileCashieringDetailType
 */
export interface ProfileCashieringDetailType {
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof ProfileCashieringDetailType
     */
    paymentMethod?: CodeDescriptionType;
    /**
     *
     * @type {UniqueIDType}
     * @memberof ProfileCashieringDetailType
     */
    eCommerceId?: UniqueIDType;
    /**
     *
     * @type {ProfileRoutingInstructionsType}
     * @memberof ProfileCashieringDetailType
     */
    routingInstructions?: ProfileRoutingInstructionsType;
    /**
     * Tax type code.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    taxType?: string;
    /**
     * Guest type code.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    fiscalGuestType?: string;
    /**
     * Hotel Code for which the routing instructions are provided for a profile.
     * @type {string}
     * @memberof ProfileCashieringDetailType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the ProfileCashieringDetailType interface.
 */
export declare function instanceOfProfileCashieringDetailType(value: object): boolean;
export declare function ProfileCashieringDetailTypeFromJSON(json: any): ProfileCashieringDetailType;
export declare function ProfileCashieringDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCashieringDetailType;
export declare function ProfileCashieringDetailTypeToJSON(value?: ProfileCashieringDetailType | null): any;
