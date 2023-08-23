/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Membership Type Rules reservation conditions.
 * @export
 * @interface MembershipTypeRuleReservationConditionsType
 */
export interface MembershipTypeRuleReservationConditionsType {
    /**
     * Reservations with start date for which the rule is applied.
     * @type {Date}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    reservationStartDate?: Date;
    /**
     * Reservations with end date for which the rule is applied.
     * @type {Date}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    reservationEndDate?: Date;
    /**
     * Minimum length of stay required to receive points if membership type is based on STAY.
     * @type {number}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    minimumNights?: number;
    /**
     * Maximum length of stay required to receive points if membership type is based on NIGHTS.
     * @type {number}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    maximumNights?: number;
    /**
     * Number of days from enrolment.
     * @type {number}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    daysFromEnrollment?: number;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    sunday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    monday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    tuesday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    wednesday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    thursday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    friday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof MembershipTypeRuleReservationConditionsType
     */
    saturday?: boolean;
}
/**
 * Check if a given object implements the MembershipTypeRuleReservationConditionsType interface.
 */
export declare function instanceOfMembershipTypeRuleReservationConditionsType(value: object): boolean;
export declare function MembershipTypeRuleReservationConditionsTypeFromJSON(json: any): MembershipTypeRuleReservationConditionsType;
export declare function MembershipTypeRuleReservationConditionsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTypeRuleReservationConditionsType;
export declare function MembershipTypeRuleReservationConditionsTypeToJSON(value?: MembershipTypeRuleReservationConditionsType | null): any;
